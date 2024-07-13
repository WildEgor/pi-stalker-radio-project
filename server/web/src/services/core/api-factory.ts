import HealthRPCClient from "@/services/api/health-rpc.ts";
import HttpService from "@/services/core/base-http.ts";

const ApiList = [
    {
        variableName: "HealthAPI",
        classEntity: HealthRPCClient,
    },
];

export interface ApiFactoryParams {
    httpService: HttpService;
}

interface ApiFactory {
    healthRPC: HealthRPCClient;
}

class ApiFactory {
    protected httpService: HttpService;

    [index: string]: any;

    constructor({ httpService }: ApiFactoryParams) {
        this.httpService = httpService;

        ApiList.forEach((api) => {
            this[api.variableName] = new api.classEntity({
                httpService,
            });
        });
    }
}

export default ApiFactory;
