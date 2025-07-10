import type { ResourceProvider } from '../interfaces/resourcePrices'

export const services: ResourceProvider[] = [
  {
    id: '1',
    Service: 'Transfer in/out Airport (day)',
    type: 'Guide',
    VehicleType: '',
    refPrices: [
      {
        minPersons: 1,
        maxPersons: 2,
        price: 152,
        currency: 'BOL',
        isPerPerson: false,
      },
      {
        minPersons: 3,
        maxPersons: 5,
        price: 152,
        currency: 'BOL',
        isPerPerson: false,
      },
      {
        minPersons: 6,
        maxPersons: 8,
        price: 152,
        currency: 'BOL',
        isPerPerson: false,
      },
    ],
  },
  {
    id: '2',
    Service: 'Transfer in/out Airport (night)',
    type: 'Guide',
    VehicleType: '',
    refPrices: [
      {
        minPersons: 1,
        maxPersons: 2,
        price: 152,
        currency: 'BOL',
        isPerPerson: false,
      },
      {
        minPersons: 3,
        maxPersons: 5,
        price: 152,
        currency: 'BOL',
        isPerPerson: false,
      },
      {
        minPersons: 6,
        maxPersons: 8,
        price: 152,
        currency: 'BOL',
        isPerPerson: false,
      },
    ],
  },
  {
    id: '3',
    Service: 'Transfer in/out Airport (day)',
    type: 'Transport',
    VehicleType: 'Car',
    refPrices: [
      {
        minPersons: 1,
        maxPersons: 3,
        price: 152,
        currency: 'BOL',
        isPerPerson: false,
      },
      {
        minPersons: 4,
        maxPersons: 4,
        price: 152,
        currency: 'BOL',
        isPerPerson: false,
      },
    ],
  },
]
