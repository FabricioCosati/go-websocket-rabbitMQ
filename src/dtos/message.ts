import { UserDto } from "./user"

export interface MessageDto {
    User: UserDto
    Message: string
    Time?: string
}