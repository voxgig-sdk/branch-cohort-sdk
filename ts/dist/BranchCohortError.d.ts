import { Context } from './Context';
declare class BranchCohortError extends Error {
    isBranchCohortError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    status: number;
    get notFound(): boolean;
    constructor(code: string, msg: string, ctx: Context);
}
export { BranchCohortError };
