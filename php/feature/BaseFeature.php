<?php
declare(strict_types=1);

// BranchCohort SDK base feature

class BranchCohortBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(BranchCohortContext $ctx, array $options): void {}
    public function PostConstruct(BranchCohortContext $ctx): void {}
    public function PostConstructEntity(BranchCohortContext $ctx): void {}
    public function SetData(BranchCohortContext $ctx): void {}
    public function GetData(BranchCohortContext $ctx): void {}
    public function GetMatch(BranchCohortContext $ctx): void {}
    public function SetMatch(BranchCohortContext $ctx): void {}
    public function PrePoint(BranchCohortContext $ctx): void {}
    public function PreSpec(BranchCohortContext $ctx): void {}
    public function PreRequest(BranchCohortContext $ctx): void {}
    public function PreResponse(BranchCohortContext $ctx): void {}
    public function PreResult(BranchCohortContext $ctx): void {}
    public function PreDone(BranchCohortContext $ctx): void {}
    public function PreUnexpected(BranchCohortContext $ctx): void {}
}
