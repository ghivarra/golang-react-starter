export interface APIResponse {
    status: 'error' | 'success';
    message: string;
    data?: unknown;
    errors?: ErrorResponse;
}

export interface ErrorResponse {
    [key: string]: string[];
}

export interface UserData {
    ID: number;
    Name: string;
    Username: string;
    Email: string;
    Password: string;
    IsSuperadmin: number;
    RoleID: number;
    RoleName: string;
    CreatedAt: string;
    UpdatedAt: string;
}