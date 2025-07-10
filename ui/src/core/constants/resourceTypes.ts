import { colors } from './colors'

export const ResourceTypes = [
  {
    label: 'Service',
    value: 'service',
    color: colors.yellow,
  },
  {
    label: 'Transport',
    value: 'transport',
    color: colors.red,
  },
  {
    label: 'Activity',
    value: 'activity',
    color: colors.purple,
  },
  {
    label: 'Hotel',
    value: 'hotel',
    color: colors.green,
  },
  {
    label: 'Simple',
    value: 'simple',
    color: colors.green,
  },
  {
    label: 'Double',
    value: 'double',
    color: colors.green,
  },
  {
    label: 'Shared',
    value: 'shared',
    color: colors.green,
  },
  {
    label: 'Breackfast',
    value: 'breackfast',
    color: colors.red,
  },
  {
    label: 'Lunch',
    value: 'lunch',
    color: colors.red,
  },
  {
    label: 'Dinner',
    value: 'dinner',
    color: colors.red,
  },
  {
    label: 'Extra',
    value: 'extra',
    color: colors.red,
  },
] as const
