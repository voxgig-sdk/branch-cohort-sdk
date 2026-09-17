# BranchCohort SDK exists test

import pytest
from branchcohort_sdk import BranchCohortSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = BranchCohortSDK.test(None, None)
        assert testsdk is not None
