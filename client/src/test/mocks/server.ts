import { setupServer } from 'msw/node';
import { http, HttpResponse } from 'msw';

export const server = setupServer(
  http.post('http://localhost:8080/api/contributions', async ({ request }) => {
    const body = await request.json();
    return HttpResponse.json({
      id: 999,
      title: body?.title,
      repository: body?.repository,
      date: body?.date,
      url: body?.url,
    }, { status: 201 });
  }),

  http.get('http://localhost:8080/api/contributions', ({ request }) => {
    const url = new URL(request.url);
    const repo = url.searchParams.get('repository');

    if (repo) {
      return HttpResponse.json([
        {
          id: 1,
          title: "Test Contribution",
          repository: repo,
          date: "2026-04-17",
          url: "https://github.com/pull/1"
        }
      ]);
    }

    return HttpResponse.json([]);
  })
);