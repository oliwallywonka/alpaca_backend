import Pocketbase from 'pocketbase'
import { BASE_URL } from '@/core/settings'

const pb = new Pocketbase(BASE_URL)
export { pb as API }
