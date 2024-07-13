import HttpService from "@/services/core/base-http.ts";

export interface BaseRPCRequest<T> {
    method: string;
    params?: T[];
    id: number;
}

export interface BaseRPCResponse<T> {
    result: T;
    error: string;
    id: number;
}

interface BaseRPCServiceParams {
    name: string;
    httpService: HttpService;
}

class BaseRpc {
    static rpcName: string = "name not implemented";

    protected httpService: HttpService;
    private readonly name: string;

    constructor({ name, httpService }: BaseRPCServiceParams) {
        this.name = name;
        this.httpService = httpService;
    }

    protected get rpc() {
        return `${this.name}`;
    }

    public async call<TP = [], TR = void>({method, params, id}: BaseRPCRequest<TP>): Promise<BaseRPCResponse<TR> | void> {
        try {
            const resp = await this.httpService.post<BaseRPCRequest<TP>, BaseRPCResponse<TR>>("/rpc", {
                method: `${this.name}.${method}`,
                params,
                id,
            })

            return resp
        } catch (e: unknown) {
            console.error(e)
        }

        return
    }
}

export default BaseRpc;
