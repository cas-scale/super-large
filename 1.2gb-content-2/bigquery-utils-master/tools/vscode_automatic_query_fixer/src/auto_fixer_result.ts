export interface QueryFix {
  status: string;
  error?: string;
  approach?: string;
  errorPosition?: ErrorPosition;
  options?: FixOption[];
}

export interface ErrorPosition {
  row: number;
  column: number;
}

export interface FixOption {
  action: string;
  fixedQuery: string;
}
// ID-1768294462-b9dabd88
