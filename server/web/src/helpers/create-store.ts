import ApiFactory from "@/services/core/api-factory.ts";
import HttpService from "@/services/core/base-http.ts";

export interface Environment {
    rpc_api: string;
}

export interface RootEnv {
    apiFactory: ApiFactory;
    envConfig: Environment;
}

export interface CreateStoreResult {
    env: RootEnv;

}

export interface CreateStoreOptions {
    envConfig: Environment;
}

const createStore = ({ envConfig }: CreateStoreOptions): CreateStoreResult => {
    const httpService = new HttpService({ baseURL: envConfig.rpc_api });
    const apiFactory = new ApiFactory({ httpService });

    const env: RootEnv = {
        apiFactory,
        envConfig,
    };

    return {
        env,
    };
};

export default createStore;
