export namespace collections {
	
	export class Request {
	    name: string;
	    description: string;
	    url: string;
	    method: string;
	    body: string;
	    query_params: string;
	    path_params: string;
	
	    static createFrom(source: any = {}) {
	        return new Request(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	        this.url = source["url"];
	        this.method = source["method"];
	        this.body = source["body"];
	        this.query_params = source["query_params"];
	        this.path_params = source["path_params"];
	    }
	}
	export class Folder {
	    name: string;
	    description: string;
	    items: CollectionItem[];
	
	    static createFrom(source: any = {}) {
	        return new Folder(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	        this.items = this.convertValues(source["items"], CollectionItem);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class CollectionItem {
	    is_folder: boolean;
	    folder?: Folder;
	    request?: Request;
	
	    static createFrom(source: any = {}) {
	        return new CollectionItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.is_folder = source["is_folder"];
	        this.folder = this.convertValues(source["folder"], Folder);
	        this.request = this.convertValues(source["request"], Request);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Collection {
	    name: string;
	    description: string;
	    variables: environments.EnvironmentItem[];
	    environments: environments.Environment[];
	    environment: environments.Environment;
	    items: CollectionItem[];
	
	    static createFrom(source: any = {}) {
	        return new Collection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	        this.variables = this.convertValues(source["variables"], environments.EnvironmentItem);
	        this.environments = this.convertValues(source["environments"], environments.Environment);
	        this.environment = this.convertValues(source["environment"], environments.Environment);
	        this.items = this.convertValues(source["items"], CollectionItem);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	

}

export namespace environments {
	
	export class EnvironmentItem {
	    name: string;
	    value: string;
	    initial_value: string;
	
	    static createFrom(source: any = {}) {
	        return new EnvironmentItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.value = source["value"];
	        this.initial_value = source["initial_value"];
	    }
	}
	export class Environment {
	    name: string;
	    description: string;
	    items: EnvironmentItem[];
	
	    static createFrom(source: any = {}) {
	        return new Environment(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	        this.items = this.convertValues(source["items"], EnvironmentItem);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

