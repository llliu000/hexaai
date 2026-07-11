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
import {
  Button,
  Empty,
  Input,
  InputNumber,
  Modal,
  Typography,
} from '@douyinfe/semi-ui';
import { IconDelete, IconPlus } from '@douyinfe/semi-icons';
import { API, showError, showSuccess } from '../../../../helpers';

const { Text } = Typography;

const createRow = (model = '', discount = 10000) => ({
  id: `${Date.now()}-${Math.random()}`,
  model,
  discount,
});

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

const modelDiscountsToRows = (user) => {
  return Object.entries(parseModelDiscounts(user)).map(([model, discount]) =>
    createRow(model, discount),
  );
};

const normalizeModelDiscounts = (rows) => {
  const normalized = {};
  for (const row of rows) {
    const model = row.model.trim();
    const discount = Number(row.discount);
    if (!model || !Number.isInteger(discount) || discount <= 0) {
      return null;
    }
    if (normalized[model] !== undefined) {
      return null;
    }
    normalized[model] = discount;
  }
  return normalized;
};

const UserModelDiscountsModal = ({ visible, onCancel, user, t, onSuccess }) => {
  const [rows, setRows] = useState([]);
  const [saving, setSaving] = useState(false);

  useEffect(() => {
    if (visible) {
      setRows(modelDiscountsToRows(user));
    }
  }, [visible, user]);

  const addRow = () => {
    setRows((currentRows) => [...currentRows, createRow()]);
  };

  const updateRow = (id, field, value) => {
    setRows((currentRows) =>
      currentRows.map((row) =>
        row.id === id ? { ...row, [field]: value } : row,
      ),
    );
  };

  const removeRow = (id) => {
    setRows((currentRows) => currentRows.filter((row) => row.id !== id));
  };

  const handleSave = async () => {
    if (!user?.id) {
      return;
    }
    const modelDiscounts = normalizeModelDiscounts(rows);
    if (!modelDiscounts) {
      showError(t('请填写不重复的模型名称，折扣值必须是正整数'));
      return;
    }

    setSaving(true);
    try {
      const res = await API.put(`/api/user/${user.id}/model_discounts`, {
        model_discounts: modelDiscounts,
      });
      const { success, message } = res.data;
      if (success) {
        showSuccess(t('折扣配置已更新'));
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
      title={t('折扣配置')}
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
        <div className='flex items-center justify-between gap-3'>
          <Text type='tertiary' size='small'>
            {t('10000 表示 1.0，例如 8000 表示 0.8 倍价格。')}
          </Text>
          <Button
            icon={<IconPlus />}
            theme='solid'
            type='primary'
            onClick={addRow}
          >
            {t('新增')}
          </Button>
        </div>
        {rows.length === 0 ? (
          <Empty description={t('暂无折扣配置')} style={{ padding: '32px 0' }} />
        ) : (
          <div className='space-y-2'>
            <div className='grid grid-cols-[1fr_150px_36px] gap-2 text-xs px-1'>
              <Text type='tertiary'>{t('模型名称')}</Text>
              <Text type='tertiary'>{t('折扣值')}</Text>
              <span />
            </div>
            {rows.map((row) => (
              <div
                key={row.id}
                className='grid grid-cols-[1fr_150px_36px] gap-2 items-center'
              >
                <Input
                  value={row.model}
                  placeholder='gpt-4o'
                  onChange={(value) => updateRow(row.id, 'model', value)}
                />
                <InputNumber
                  value={row.discount}
                  min={1}
                  step={1}
                  precision={0}
                  placeholder='10000'
                  onChange={(value) =>
                    updateRow(
                      row.id,
                      'discount',
                      value === '' || value == null ? '' : value,
                    )
                  }
                  style={{ width: '100%' }}
                />
                <Button
                  theme='borderless'
                  type='danger'
                  icon={<IconDelete />}
                  onClick={() => removeRow(row.id)}
                />
              </div>
            ))}
          </div>
        )}
        <Text type='tertiary' size='small'>
          {t('删除所有行并保存即可清空配置。')}
        </Text>
      </div>
    </Modal>
  );
};

export default UserModelDiscountsModal;
