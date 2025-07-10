import { ref } from 'vue'
import type { ParamsRequest } from '../interfaces/api'

export function useParams(paramParams: Partial<ParamsRequest> = {}) {
  const defaultParams: ParamsRequest = {
    page: 1,
    perPage: 10,
    filter: '',
    orderBy: 'created',
    orderDirection: 'DESC',
    ...paramParams,
  }
  const params = ref<ParamsRequest>(defaultParams)

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

  const setParams = (newParams: Partial<ParamsRequest>) => {
    params.value = { ...params.value, ...newParams }
  }

  return {
    params,
    setParams,
    resetParams,
    getUrlSearchParams,
  }
}
