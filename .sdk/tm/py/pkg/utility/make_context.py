# BranchCohort SDK utility: make_context

from projectname_sdk.core.context import BranchCohortContext


def make_context_util(ctxmap, basectx):
    return BranchCohortContext(ctxmap, basectx)
