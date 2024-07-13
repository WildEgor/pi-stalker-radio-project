import axios, { AxiosInstance } from "axios";

export interface HttpServiceConfig {
    baseURL?: string;
}
export default class HttpService {
    private axiosInstance: AxiosInstance;

    constructor({ baseURL }: HttpServiceConfig) {
        this.axiosInstance = axios.create({
            baseURL,
            headers: {
                "content-type": "application/json",
                "cache-control": "no-cache"
            },
        });
    }

    async get<T>(url: string): Promise<T> {
        const response = await this.axiosInstance.get(url);
        return response.data as T;
    }

    async post<T, R>(url: string, data?: T) {
        const response = await this.axiosInstance.post(url, data);
        return response.data as R;
    }

    async put<T, R>(url: string, data: T) {
        const response = await this.axiosInstance.put(url, data);
        return response.data as R;
    }

    async delete<R>(url: string) {
        const response = await this.axiosInstance.delete(url);
        return response.data as R;
    }
}
