import { ContributionRepository } from '../repositories/contribution-repo.js';
import { Contribution, CreateContribution } from '../types/contribution.js';

export class ContributionService {
  constructor(private readonly repository: ContributionRepository) {}

  async logContribution(
    title: string,
    repository: string,
    url?: string,
    date: string = new Date().toISOString().split('T')[0]
  ): Promise<Contribution> {
    const payload: CreateContribution = { title, repository, date, url };
    return this.repository.add(payload);
  }

  async getContributions(repository?: string): Promise<Contribution[]> {
    return this.repository.list(repository);
  }
}