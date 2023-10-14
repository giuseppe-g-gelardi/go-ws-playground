
import { type ValueOf } from "~/types";


const baseUrl = process.env.BASE_API_URL;

type UrlReturnType = `${typeof baseUrl}/${ValueOf<typeof Collections>}${string | null}`

export function URL(collection: ValueOf<typeof Collections>): UrlReturnType;
export function URL(collection: ValueOf<typeof Collections>, id: string): UrlReturnType
export function URL(collection: ValueOf<typeof Collections>, id?: string): UrlReturnType {
  // return `${baseUrl}/${collection}${id ? `/${id}` : null}}`;
  if (!id) {
    return `${baseUrl}/${collection}`;
  }
  return `${baseUrl}/${collection}/${id}`;
}


const Collections = {
  user: 'user',
  users: 'users',
} as const


