import { v4 as uuidv4 } from "uuid";

export const uniqueId = (length: number): string => {
  let result = ""
  const chars = uuidv4().replaceAll(/-/gi, "")
  let charsLength = chars.length
  let counter = 0

  if (length >= charsLength) {
    charsLength = length
  }

  while (counter < length) {
    result += chars.charAt(Math.floor(Math.random() * charsLength))
    counter += 1
  }

  return result
}
