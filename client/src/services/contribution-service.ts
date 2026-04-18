
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
    if (!title || title.trim() === '') {
      throw new Error('Title is required and cannot be empty');
    }

    if (!repository || repository.trim() === '') {
      throw new Error('Repository is required and cannot be empty');
    }

    const payload: CreateContribution = {
      title: title.trim(),
      repository: repository.trim(),
      date,
      url: url?.trim(),
    };

    return this.repository.add(payload);
  }

  async getContributions(repository?: string): Promise<Contribution[]> {
    const contributions = await this.repository.list(repository);
    if (!contributions || !Array.isArray(contributions)) {
      return [];
    }

    return contributions;
  }
}