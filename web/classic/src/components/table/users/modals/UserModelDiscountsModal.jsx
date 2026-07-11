/*
Copyright (C) 2025 QuantumNous

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.

For commercial licensing, please contact support@quantumnous.com
*/

import React, { useEffect, useState } from 'react';
import { Modal, TextArea, Typography } from '@douyinfe/semi-ui';
import { API, showError, showSuccess } from '../../../../helpers';

const { Text } = Typography;

const parseModelDiscounts = (user) => {
  if (!user?.setting) {
    return {};
  }
  try {
    const setting = JSON.parse(user.setting);
    if (
      setting?.model_discounts &&
      typeof setting.model_discounts === 'object' &&
      !Array.isArray(setting.model_discounts)
    ) {
      return setting.model_discounts;
    }
  } catch (e) {
    return {};
  }
  return {};
};

const normalizeModelDiscounts = (value) => {
  const trimmed = value.trim();
  if (!trimmed) {
    return {};
  }

  let parsed;
  try {
    parsed = JSON.parse(trimmed);
  } catch (e) {
    return null;
  }

  if (!parsed || typeof parsed !== 'object' || Array.isArray(parsed)) {
    return null;
  }

  const normalized = {};
  for (const [modelName, discount] of Object.entries(parsed)) {
    const key = modelName.trim();
    if (
      !key ||
      typeof discount !== 'number' ||
      !Number.isInteger(discount) ||
      discount <= 0
    ) {
      return null;
    }
    normalized[key] = discount;
  }
  return normalized;
};

const UserModelDiscountsModal = ({ visible, onCancel, user, t, onSuccess }) => {
  const [value, setValue] = useState('{}');
  const [saving, setSaving] = useState(false);

  useEffect(() => {
    if (visible) {
      setValue(JSON.stringify(parseModelDiscounts(user), null, 2));
    }
  }, [visible, user]);

  const handleSave = async () => {
    if (!user?.id) {
      return;
    }
    const modelDiscounts = normalizeModelDiscounts(value);
    if (!modelDiscounts) {
      showError(t('模型价格配置必须是 JSON 对象，且值为正整数'));
      return;
    }

    setSaving(true);
    try {
      const res = await API.put(`/api/user/${user.id}/model_discounts`, {
        model_discounts: modelDiscounts,
      });
      const { success, message } = res.data;
      if (success) {
        showSuccess(t('模型价格配置已更新'));
        onSuccess?.();
        onCancel?.();
      } else {
        showError(message || t('保存失败'));
      }
    } catch (e) {
      showError(t('请求失败'));
    } finally {
      setSaving(false);
    }
  };

  return (
    <Modal
      title={t('模型价格配置')}
      visible={visible}
      onCancel={onCancel}
      onOk={handleSave}
      confirmLoading={saving}
      okText={t('保存')}
      cancelText={t('取消')}
      width={640}
    >
      <div className='space-y-3'>
        {user?.username ? (
          <Text type='tertiary'>
            {t('目标用户：{{username}}', { username: user.username })}
          </Text>
        ) : null}
        <TextArea
          value={value}
          onChange={setValue}
          autosize={{ minRows: 10, maxRows: 16 }}
          placeholder='{"gpt-4o": 8000}'
          style={{ fontFamily: 'monospace' }}
        />
        <Text type='tertiary' size='small'>
          {t('10000 表示 1.0，例如 8000 表示 0.8 倍价格。留空或填写 {} 可清空配置。')}
        </Text>
      </div>
    </Modal>
  );
};

export default UserModelDiscountsModal;
