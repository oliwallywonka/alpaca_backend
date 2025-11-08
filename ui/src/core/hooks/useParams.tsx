import { ref } from 'vue'
import type { QueryParams } from '../interfaces/api'

export function useParams(paramParams: Partial<QueryParams> = {}) {
  const defaultParams: QueryParams = {
    page: 1,
    perPage: 10,
    filter: '',
    orderBy: 'created',
    orderDirection: 'DESC',
    expand: '',
    ...paramParams,
  }
  const params = ref<QueryParams>(defaultParams)

  const resetParams = () => {
    params.value = defaultParams
  }

  const getUrlSearchParams = () => {
    const urlParams = new URLSearchParams()
    Object.entries(params.value).forEach(([key, value]) => {
      urlParams.set(key, String(value))
    })
    return urlParams.toString()
  }

  const setParams = (newParams: Partial<QueryParams>) => {
    params.value = { ...params.value, ...newParams }
  }

  return {
    params,
    setParams,
    resetParams,
    getUrlSearchParams,
  }
}
