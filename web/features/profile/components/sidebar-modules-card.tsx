/*
Copyright (C) 2023-2026 QuantumNous

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
import { Ellipsis } from 'lucide-react'
import { useCallback, useEffect, useRef, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { toast } from 'sonner'

import { Button } from '@/components/ui/button'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import { Switch } from '@/components/ui/switch'
import { api } from '@/lib/api'
import { useAuthStore } from '@/stores/auth-store'

import { DesignIcon } from './design-icon'

type SidebarModuleConfig = {
  enabled: boolean
  [key: string]: boolean
}

type SidebarModulesConfig = Record<string, SidebarModuleConfig>

type SectionDef = {
  key: string
  title: string
  description: string
  modules: { key: string; title: string; description: string }[]
}

export function SidebarModulesCard() {
  const { t } = useTranslation()
  const [loading, setLoading] = useState(false)
  const [dirty, setDirty] = useState(false)
  const [advanced, setAdvanced] = useState(false)
  const [config, setConfig] = useState<SidebarModulesConfig>({})
  const savedConfig = useRef<SidebarModulesConfig>({})
  const currentUser = useAuthStore((s) => s.auth.user)
  const setUser = useAuthStore((s) => s.auth.setUser)

  const sectionDefs: SectionDef[] = [
    {
      key: 'chat',
      title: t('Chat Area'),
      description: t('profile.design.chatDescription', {
        defaultValue: '操控按钮和聊天功能',
      }),
      modules: [
        {
          key: 'playground',
          title: t('Playground'),
          description: t('AI model testing environment'),
        },
        {
          key: 'chat',
          title: t('Chat'),
          description: t('Chat session management'),
        },
      ],
    },
    {
      key: 'console',
      title: t('Console Area'),
      description: t('Data management and log viewing'),
      modules: [
        {
          key: 'detail',
          title: t('Dashboard'),
          description: t('System data statistics'),
        },
        {
          key: 'token',
          title: t('Token Management'),
          description: t('API token management'),
        },
        {
          key: 'log',
          title: t('Usage Logs'),
          description: t('API usage records'),
        },
        {
          key: 'midjourney',
          title: t('Drawing Logs'),
          description: t('Drawing task records'),
        },
        {
          key: 'task',
          title: t('Task Logs'),
          description: t('System task records'),
        },
      ],
    },
    {
      key: 'personal',
      title: t('Personal Center Area'),
      description: t('User personal functions'),
      modules: [
        {
          key: 'topup',
          title: t('Billing Center'),
          description: t('Balance and top-up management'),
        },
        {
          key: 'personal',
          title: t('Personal Settings'),
          description: t('Personal info settings'),
        },
      ],
    },
  ]

  const loadConfig = useCallback(async () => {
    try {
      const res = await api.get('/api/user/self')
      if (res.data.success && res.data.data?.sidebar_modules) {
        const raw = res.data.data.sidebar_modules
        const parsed = typeof raw === 'string' ? JSON.parse(raw) : raw
        savedConfig.current = parsed
        setConfig(parsed)
      } else {
        const defaults: SidebarModulesConfig = {}
        for (const sec of sectionDefs) {
          defaults[sec.key] = { enabled: true }
          for (const mod of sec.modules) defaults[sec.key][mod.key] = true
        }
        savedConfig.current = defaults
        setConfig(defaults)
      }
    } catch {
      /* ignore */
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  useEffect(() => {
    loadConfig()
  }, [loadConfig])

  const toggleSection = (sectionKey: string, val: boolean) => {
    setDirty(true)
    setConfig((prev) => ({
      ...prev,
      [sectionKey]: { ...prev[sectionKey], enabled: val },
    }))
  }

  const toggleModule = (
    sectionKey: string,
    moduleKey: string,
    val: boolean
  ) => {
    setDirty(true)
    setConfig((prev) => ({
      ...prev,
      [sectionKey]: { ...prev[sectionKey], [moduleKey]: val },
    }))
  }

  const handleSave = async () => {
    setLoading(true)
    try {
      const serialized = JSON.stringify(config)
      const res = await api.put('/api/user/self', {
        sidebar_modules: serialized,
      })
      if (res.data.success) {
        // Sync to auth-store so useSidebarConfig re-runs and the sidebar
        // updates immediately without needing a page refresh.
        if (currentUser) {
          setUser({ ...currentUser, sidebar_modules: serialized })
        }
        toast.success(t('Saved successfully'))
        savedConfig.current = config
        setDirty(false)
      } else {
        toast.error(res.data.message || t('Save failed'))
      }
    } catch {
      toast.error(t('Save failed, please retry'))
    } finally {
      setLoading(false)
    }
  }

  const handleReset = () => {
    const defaults: SidebarModulesConfig = {}
    for (const sec of sectionDefs) {
      defaults[sec.key] = { enabled: true }
      for (const mod of sec.modules) defaults[sec.key][mod.key] = true
    }
    setConfig(defaults)
    setDirty(true)
    toast.success(t('Reset to default configuration'))
  }

  const renderSection = (section: SectionDef) => (
    <section
      className={`profile-module-group group-${section.key}`}
      key={section.key}
    >
      <h3>{section.title}</h3>
      <p>{section.description}</p>
      {section.modules.map((mod) => (
        <div className='profile-module-row' key={mod.key}>
          <div>
            <strong>{mod.title}</strong>
            <small>{mod.description}</small>
          </div>
          <Switch
            aria-label={mod.title}
            checked={config[section.key]?.[mod.key] !== false}
            disabled={loading || config[section.key]?.enabled === false}
            onCheckedChange={(value) =>
              toggleModule(section.key, mod.key, value)
            }
          />
        </div>
      ))}
    </section>
  )

  return (
    <>
      <Card data-card-hover='false' className='profile-modules-main'>
        <CardHeader>
          <DesignIcon name='modules' />
          <div>
            <CardTitle>{t('Sidebar Personal Settings')}</CardTitle>
            <CardDescription>
              {t('Customize sidebar display content')}
            </CardDescription>
          </div>
          <button
            className='profile-module-more'
            type='button'
            title={t('Settings')}
            aria-label={t('Sidebar Settings')}
            onClick={() => setAdvanced(true)}
          >
            <Ellipsis size={16} />
          </button>
        </CardHeader>
        <CardContent>{sectionDefs.slice(0, 2).map(renderSection)}</CardContent>
      </Card>
      <Card data-card-hover='false' className='profile-modules-personal'>
        <CardContent>{renderSection(sectionDefs[2])}</CardContent>
      </Card>
      {dirty && (
        <div className='profile-module-save' role='status'>
          <Button
            variant='outline'
            onClick={() => {
              setConfig(savedConfig.current)
              setDirty(false)
            }}
            disabled={loading}
          >
            {t('Cancel')}
          </Button>
          <Button onClick={handleSave} disabled={loading}>
            {loading ? t('Saving...') : t('Save Changes')}
          </Button>
        </div>
      )}
      <Dialog open={advanced} onOpenChange={setAdvanced}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>{t('Sidebar Personal Settings')}</DialogTitle>
          </DialogHeader>
          {sectionDefs.map((section) => (
            <div
              className='flex items-center justify-between gap-4'
              key={section.key}
            >
              <span>{section.title}</span>
              <Switch
                aria-label={section.title}
                checked={config[section.key]?.enabled !== false}
                disabled={loading}
                onCheckedChange={(value) => toggleSection(section.key, value)}
              />
            </div>
          ))}
          <Button variant='outline' onClick={handleReset}>
            {t('Reset to Default')}
          </Button>
          <Button
            onClick={async () => {
              await handleSave()
            }}
            disabled={loading}
          >
            {loading ? t('Saving...') : t('Save Changes')}
          </Button>
        </DialogContent>
      </Dialog>
    </>
  )
}
