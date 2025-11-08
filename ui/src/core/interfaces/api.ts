export interface ApiResponse<T> {
  data: T
  message: string
  error: string[]
}

export interface Paginated<T> {
  total: number
  items: T[]
  page: number
}

export interface QueryParams {
  page: number
  perPage: number
  filter: string
  orderBy: string
  orderDirection: 'ASC' | 'DESC'
  expand: string
}
