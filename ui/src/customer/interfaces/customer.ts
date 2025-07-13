import type { ContactField } from "@/core/interfaces/fields"

export type Title = 'Mr' | 'Mrs' | 'Miss' | 'Ms' | 'Dr'

export interface Customer {
    id: string
    title: Title
    firstName: string
    middleName?: string
    lastName: string
    dateOfBirth: string
    contacts: ContactField[]
    isLead: boolean
    created: string
    updated: string
}