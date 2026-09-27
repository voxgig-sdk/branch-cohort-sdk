"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.BranchCohortError = void 0;
class BranchCohortError extends Error {
    isBranchCohortError = true;
    sdk = 'BranchCohort';
    code;
    ctx;
    status = -1;
    // `err.notFound` rather than a magic number at every call site.
    get notFound() { return 404 === this.status; }
    constructor(code, msg, ctx) {
        super(msg);
        this.code = code;
        this.ctx = ctx;
    }
}
exports.BranchCohortError = BranchCohortError;
//# sourceMappingURL=BranchCohortError.js.map