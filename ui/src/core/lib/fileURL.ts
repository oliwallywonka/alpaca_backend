import { BASE_URL } from '@/core/settings'

interface Data {
  collectionName: string
  recordID: string
  fileName: string
}

export function getFileURL({ collectionName, recordID, fileName }: Data) {
  return `${BASE_URL}/api/files/${collectionName}/${recordID}/${fileName}`
}

export function getFileNameFromURL(url: string) {
  const urlParts = url.split('/')
  return urlParts[urlParts.length - 1]
}
