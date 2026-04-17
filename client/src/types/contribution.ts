export interface Contribution {
  id?: number;
  title: string;
  repository: string;
  date: string;
  url?: string;
}

export type CreateContribution = Omit<Contribution, 'id'>;