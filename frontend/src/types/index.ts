export interface APIResponse {
    status: 'error' | 'success';
    message: string;
    data?: unknown;
    errors?: unknown;
}