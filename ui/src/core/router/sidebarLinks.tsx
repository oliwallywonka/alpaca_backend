import {
  AudioWaveform,
  Banknote,
  BarChart2,
  Command,
  GalleryVerticalEnd,
  Map,
  MapPin,
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
      name: 'Alpaca Face',
      logo: GalleryVerticalEnd,
      plan: 'Travel Agency',
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
      title: 'People Management',
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
        {
          title: 'Customers',
          url: '/customers',
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
          title: 'Tours Templates',
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
      ],
    },

    {
      title: 'Trips Management',
      url: '/trips',
      icon: MapPin,
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
