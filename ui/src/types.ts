export interface IApplication {
    id: number;
    token: string;
    name: string;
    description: string;
    image: string;
    internal: boolean;
}

export interface IClient {
    id: number;
    token: string;
    name: string;
}

export interface IPlugin {
    id: number;
    token: string;
    name: string;
    modulePath: string;
    enabled: boolean;
    author?: string;
    website?: string;
    license?: string;
    capabilities: Array<'webhooker' | 'displayer' | 'configurer' | 'messenger' | 'storager'>;
}

export interface IMessage {
    id: number;
    appid: number;
    message: string;
    title: string;
    priority: number;
    date: string;
    image?: string;
    extras?: IMessageExtras;
}

export interface IMessageExtras {
    [key: string]: any; // tslint:disable-line no-any
}

export interface IPagedMessages {
    paging: IPaging;
    messages: IMessage[];
}

export interface IPaging {
    next?: string;
    since?: number;
    size: number;
    limit: number;
}

export interface IUser {
    id: number;
    username: string;
    admin: boolean;
}

export interface IVersion {
    version: string;
    commit: string;
    buildDate: string;
}


