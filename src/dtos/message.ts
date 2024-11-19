import { UserDto } from "./user"

export interface MessageDto {
    user: UserDto
    message: string
}