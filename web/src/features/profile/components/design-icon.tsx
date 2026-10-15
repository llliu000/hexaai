import {
  Activity,
  Box,
  ChartNoAxesCombined,
  ChevronDown,
  ChevronRight,
  CircleHelp,
  CreditCard,
  Fingerprint,
  Globe2,
  KeyRound,
  Laptop,
  LayoutGrid,
  Link2,
  LogOut,
  Mail,
  MonitorSmartphone,
  Radio,
  Search,
  ServerCog,
  Settings,
  Settings2,
  ShieldCheck,
  ShieldUser,
  SlidersHorizontal,
  Ticket,
  Trash2,
  UserRound,
  Users,
  WalletCards,
  type LucideIcon,
} from 'lucide-react'
import type { ReactElement } from 'react'

import { PageIcon, type PageIconTone } from '@/components/ui/page-icon'

import arrowRightIcon from '../assets/icons/arrow-right.svg'
import bellIcon from '../assets/icons/bell.svg'
import dashboardIcon from '../assets/icons/dashboard.svg'
import docsIcon from '../assets/icons/docs.svg'
import fileIcon from '../assets/icons/file.svg'
import keyIcon from '../assets/icons/key.svg'
import overviewIcon from '../assets/icons/overview.svg'
import tasksIcon from '../assets/icons/tasks.svg'

const designAssets = {
  overview: overviewIcon,
  dashboard: dashboardIcon,
  key: keyIcon,
  file: fileIcon,
  tasks: tasksIcon,
  docs: docsIcon,
  bell: bellIcon,
  'arrow-right': arrowRightIcon,
} as const

const sectionIcons = {
  settings: { icon: Settings2, tone: 'blue' },
  language: { icon: Globe2, tone: 'purple' },
  security: { icon: ShieldCheck, tone: 'green' },
  balance: { icon: WalletCards, tone: 'green' },
  usage: { icon: ChartNoAxesCombined, tone: 'blue' },
  requests: { icon: Activity, tone: 'purple' },
  password: { icon: KeyRound, tone: 'blue' },
  token: { icon: ShieldUser, tone: 'amber' },
  delete: { icon: Trash2, tone: 'red' },
  modules: { icon: LayoutGrid, tone: 'blue' },
  passkey: { icon: Fingerprint, tone: 'cyan' },
  mail: { icon: Mail, tone: 'blue' },
  'two-factor': { icon: ShieldCheck, tone: 'green' },
  sessions: { icon: MonitorSmartphone, tone: 'cyan' },
} satisfies Record<string, { icon: LucideIcon; tone: PageIconTone }>

const inlineIcons = {
  search: Search,
  binding: Link2,
  preferences: SlidersHorizontal,
  laptop: Laptop,
  logout: LogOut,
  'chevron-down': ChevronDown,
  'chevron-right': ChevronRight,
  wallet: WalletCards,
  user: UserRound,
  help: CircleHelp,
  channel: Radio,
  models: Box,
  users: Users,
  redemption: Ticket,
  subscriptions: CreditCard,
  'system-info': ServerCog,
  'system-settings': Settings,
} satisfies Record<string, LucideIcon>

export type DesignIconName =
  | keyof typeof designAssets
  | keyof typeof sectionIcons
  | keyof typeof inlineIcons

export function DesignIcon(props: {
  name: DesignIconName
  className?: string
}): ReactElement {
  const className = `profile-design-icon icon-${props.name} ${props.className ?? ''}`
  if (props.name in sectionIcons) {
    const definition = sectionIcons[props.name as keyof typeof sectionIcons]
    return (
      <PageIcon
        icon={definition.icon}
        tone={definition.tone}
        className={className}
      />
    )
  }
  if (props.name in inlineIcons) {
    const Icon = inlineIcons[props.name as keyof typeof inlineIcons]
    return (
      <Icon
        aria-hidden='true'
        className={className}
        size={20}
        strokeWidth={2}
        fill='none'
      />
    )
  }
  return (
    <img
      src={designAssets[props.name as keyof typeof designAssets]}
      alt=''
      aria-hidden='true'
      className={className}
      draggable={false}
    />
  )
}

export function ProfileMailIcon(props: { className?: string }) {
  return <DesignIcon name='mail' className={props.className} />
}
