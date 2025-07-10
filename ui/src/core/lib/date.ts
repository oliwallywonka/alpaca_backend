export function getStringDate(date: number) {
  return new Intl.DateTimeFormat('en-US', {
    dateStyle: 'long',
    timeStyle: 'medium',
  }).format(new Date(date * 1000))
}

export function parseDate(date: string) {
  return new Intl.DateTimeFormat('en-US', {
    dateStyle: 'long',
    timeStyle: 'medium',
  }).format(new Date(date))
}

export function addTime(date: Date, hours: number, minutes: number) {
  return new Date(date.getTime() + hours * 60 * 60 * 1000 + minutes * 60 * 1000)
}

export function substractTime(date: Date, hours: number, minutes: number) {
  return new Date(date.getTime() - hours * 60 * 60 * 1000 - minutes * 60 * 1000)
}
