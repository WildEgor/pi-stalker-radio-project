import HttpService from "@/services/core/base-http.ts";
import BaseRpc from "@/services/core/base-rpc.ts";

export interface IRPCClientProps {
    httpService: HttpService;
}

export default class HealthRPCClient extends BaseRpc {
    constructor({ httpService }: IRPCClientProps) {
        super({
            name: "HealthService",
            httpService,
        });
    }

    public async check(): Promise<void> {
        try {
           await this.call({
                method: `Check`,
                params: [],
                id: 1
            })
            console.log('check call')
        } catch (e) {
            console.log(e)
        }
    }
}
