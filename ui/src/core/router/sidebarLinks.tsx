import {
  AudioWaveform,
  Banknote,
  BarChart2,
  Command,
  GalleryVerticalEnd,
  Map,
  MapPin,
  Speech,
  Users,
} from 'lucide-vue-next'

export const sidebarData = {
  user: {
    name: 'shadcn',
    email: 'm@example.com',
    avatar: '/avatars/shadcn.jpg',
  },
  teams: [
    {
      name: 'Acme Inc',
      logo: GalleryVerticalEnd,
      plan: 'Enterprise',
    },
    {
      name: 'Acme Corp.',
      logo: AudioWaveform,
      plan: 'Startup',
    },
    {
      name: 'Evil Corp.',
      logo: Command,
      plan: 'Free',
    },
  ],
  navMain: [
    {
      title: 'Dashboard',
      url: '/',
      icon: BarChart2,
    },
    {
      title: 'User Management',
      url: '#',
      icon: Users,
      isActive: true,
      items: [
        {
          title: 'Roles',
          url: '/roles',
        },
        {
          title: 'Users',
          url: '/users',
        },
        {
          title: 'Providers',
          url: '/providers',
        },
      ],
    },
    {
      title: 'Tours Management',
      url: '#',
      icon: Map,
      isActive: true,
      items: [
        {
          title: 'Tours templates',
          url: '/tours',
        },
        {
          title: 'Tour Resources',
          url: '/resources',
        },
        {
          title: 'Destinations',
          url: '/destinations',
        },
        {
          title: 'Activities',
          url: '/activities',
        },
        {
          title: 'Meals',
          url: '/meals',
        },
        {
          title: 'Hotels',
          url: '/hotels',
        },
      ],
    },

    {
      title: 'Trips Management',
      url: '/trips',
      icon: MapPin,
    },
    {
      title: 'Tourists Management',
      url: '#',
      icon: Speech,
      isActive: true,
      items: [
        {
          title: 'Transactions',
          url: '#',
        },
      ],
    },
    {
      title: 'Cash Management',
      url: '#',
      icon: Banknote,
      isActive: true,
      items: [
        {
          title: 'Cash registers ',
          url: '/cash-registers',
        },
      ],
    },
  ],
}
