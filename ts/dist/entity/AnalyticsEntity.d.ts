import { BranchCohortEntityBase } from '../BranchCohortEntityBase';
import type { BranchCohortSDK } from '../BranchCohortSDK';
import type { Control } from '../types';
import type { Analytics, AnalyticsLoadMatch, AnalyticsCreateData } from '../BranchCohortTypes';
declare class AnalyticsEntity extends BranchCohortEntityBase<Analytics> {
    constructor(client: BranchCohortSDK, entopts: any);
    make(this: AnalyticsEntity): AnalyticsEntity;
    load(this: any, reqmatch?: AnalyticsLoadMatch, ctrl?: Control): Promise<AnalyticsEntity>;
    create(this: any, reqdata?: AnalyticsCreateData, ctrl?: Control): Promise<AnalyticsEntity>;
}
export { AnalyticsEntity };
