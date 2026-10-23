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
import type { ComponentType } from 'react'
import { useTranslation } from 'react-i18next'

import { Tabs, TabsList, TabsTrigger } from '@/components/ui/tabs'

interface ModelChartTabOption<T extends string> {
  value: T
  labelKey: string
}

interface ModelChartTabsProps<T extends string> {
  value: T
  options: readonly ModelChartTabOption<T>[]
  onValueChange: (value: T) => void
  icons?: Partial<Record<T, ComponentType<{ className?: string }>>>
}

export function ModelChartTabs<T extends string>(
  props: ModelChartTabsProps<T>
) {
  const { t } = useTranslation()

  return (
    <Tabs
      value={props.value}
      onValueChange={(value) => {
        if (value) props.onValueChange(value as T)
      }}
      className='w-full sm:w-auto'
    >
      <TabsList className='h-7 w-full overflow-visible rounded-lg border p-0.5 sm:h-8 sm:w-auto'>
        {props.options.map((option) => {
          const Icon: ComponentType<{ className?: string }> | undefined =
            props.icons?.[option.value]

          return (
            <TabsTrigger
              key={option.value}
              value={option.value}
              className='shrink-0 rounded-md px-3 text-xs font-medium transition-colors'
            >
              {Icon && <Icon className='size-3.5' />}
              {t(option.labelKey)}
            </TabsTrigger>
          )
        })}
      </TabsList>
    </Tabs>
  )
}
