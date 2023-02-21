export namespace backend {
	
	export class StIndParam {
	    indName: string;
	    timeframe: number;
	    period: number;
	    appliedPrice: number;
	    upLine: number;
	    downLine: number;
	
	    static createFrom(source: any = {}) {
	        return new StIndParam(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.indName = source["indName"];
	        this.timeframe = source["timeframe"];
	        this.period = source["period"];
	        this.appliedPrice = source["appliedPrice"];
	        this.upLine = source["upLine"];
	        this.downLine = source["downLine"];
	    }
	}
	export class StPriceRange {
	    enablePriceLimit: boolean;
	    alertWhenPriceRangeIsExceeded: boolean;
	    longMin: number;
	    longMax: number;
	    shortMin: number;
	    shortMax: number;
	
	    static createFrom(source: any = {}) {
	        return new StPriceRange(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enablePriceLimit = source["enablePriceLimit"];
	        this.alertWhenPriceRangeIsExceeded = source["alertWhenPriceRangeIsExceeded"];
	        this.longMin = source["longMin"];
	        this.longMax = source["longMax"];
	        this.shortMin = source["shortMin"];
	        this.shortMax = source["shortMax"];
	    }
	}
	export class StSocketClientSetting {
	    indParam: StIndParam;
	    priceRange: StPriceRange;
	
	    static createFrom(source: any = {}) {
	        return new StSocketClientSetting(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.indParam = this.convertValues(source["indParam"], StIndParam);
	        this.priceRange = this.convertValues(source["priceRange"], StPriceRange);
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
	export class StWsBasicInfo {
	    indNameArr: string[];
	    timeframeMap: {[key: string]: number};
	    clientStatusMap: {[key: string]: boolean};
	    appliedPriceMap: {[key: string]: number};
	    wsPort: number;
	    wsHeartbeatSeconds: number;
	    emptytStrSign: string;
	
	    static createFrom(source: any = {}) {
	        return new StWsBasicInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.indNameArr = source["indNameArr"];
	        this.timeframeMap = source["timeframeMap"];
	        this.clientStatusMap = source["clientStatusMap"];
	        this.appliedPriceMap = source["appliedPriceMap"];
	        this.wsPort = source["wsPort"];
	        this.wsHeartbeatSeconds = source["wsHeartbeatSeconds"];
	        this.emptytStrSign = source["emptytStrSign"];
	    }
	}

}

