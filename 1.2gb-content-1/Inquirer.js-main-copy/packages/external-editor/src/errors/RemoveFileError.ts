/***
 * Node External Editor
 *
 * Kevin Gravier <kevin@mrkmg.com>
 * MIT 2018
 */

export class RemoveFileError extends Error {
  originalError: Error;

  constructor(originalError: Error) {
    super(`Failed to remove temporary file. ${originalError.message}`);
    this.originalError = originalError;
  }
}
// ID-1768294475-d9c51fbc
