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
import { useQuery } from '@tanstack/react-query'
import { useMemo } from 'react'

import type { PricingData } from '@/features/pricing/types'
import { useStatus } from '@/hooks/use-status'
import { api } from '@/lib/api'
import { parseHeaderNavModulesFromStatus } from '@/lib/nav-modules'
import { useAuthStore } from '@/stores/auth-store'

export function useHomeModels() {
  const { status, loading } = useStatus()
  const user = useAuthStore((state) => state.auth.user)
  const module = parseHeaderNavModulesFromStatus(
    status as Record<string, unknown> | null
  )?.pricing
  let access: 'public' | 'login' | 'disabled' = 'public'
  if (typeof module === 'object' && !module.enabled) {
    access = 'disabled'
  } else if (typeof module === 'object' && module.requireAuth && !user) {
    access = 'login'
  }
  const query = useQuery({
    queryKey: ['homepage-models', user?.id],
    enabled: !loading && access === 'public',
    staleTime: 300000,
    queryFn: async () => {
      try {
        const response = await api.get<PricingData>('/api/pricing', {
          skipErrorHandler: true,
        })
        if (
          !response.data.success ||
          !Array.isArray(response.data.data) ||
          !Array.isArray(response.data.vendors)
        ) {
          return null
        }
        return response.data
      } catch {
        return null
      }
    },
  })
  const models = useMemo(() => {
    const vendors = new Map(
      query.data?.vendors.map((vendor) => [vendor.id, vendor]) ?? []
    )
    return (query.data?.data ?? []).map((model) => ({
      ...model,
      vendor_name: vendors.get(model.vendor_id ?? 0)?.name,
      vendor_icon: vendors.get(model.vendor_id ?? 0)?.icon,
      group_ratio: query.data?.group_ratio,
    }))
  }, [query.data])
  return {
    models,
    vendors: query.data?.vendors ?? [],
    access,
    isLoading: loading || query.isLoading,
    error: query.isFetched && !query.data,
    refetch: query.refetch,
  }
}
