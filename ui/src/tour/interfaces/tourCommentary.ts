import type { LanguageField } from "@/core/interfaces/fields"

export interface TourCommentary {
    id: string
    tour: string
    tourVariant: string
    commentary: LanguageField
    images: string[]
    startDate: string
    endDate: string
    created: string
    updated: string
}