import { describe, it, expect } from 'vitest';
import { ApiClient } from './api-client';

describe('ApiClient', () => {
  const client = new ApiClient('http://localhost:8080');

  it('should successfully add a contribution', async () => {
    const result = await client.addContribution({
      title: "Added dark mode",
      repository: "ikennarichard/contrib-tracker",
      date: "2026-04-18",
      url: "https://github.com/pull/42"
    });

    expect(result).toHaveProperty('id');
    expect(result.title).toBe("Added dark mode");
    expect(result.repository).toBe("ikennarichard/contrib-tracker");
  });

  it('should list contributions', async () => {
    const contributions = await client.listContributions();

    expect(Array.isArray(contributions)).toBe(true);
  });

  it('should list contributions with repository filter', async () => {
    const contributions = await client.listContributions("ikennarichard/contrib-tracker");

    expect(Array.isArray(contributions)).toBe(true);
    if (contributions.length > 0) {
      expect(contributions[0].repository).toBe("ikennarichard/contrib-tracker");
    }
  });
});