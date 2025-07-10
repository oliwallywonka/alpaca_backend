export interface ContactField {
  type: string
  value: string
}

export interface LanguageField {
  [key: string]: string
}

export interface LocationField {
  lat: number
  lon: number
}
