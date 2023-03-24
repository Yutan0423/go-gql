-- name: AddPullRequests :execresult
INSERT INTO pullrequests(id, base_ref_name, closed, head_ref_name, url, number, repository) VALUES
('PR_1', 'main', 1, 'feature/kinou1', 'http://example.com/repo1/pr/1', 1, 'REPO_1'),
('PR_2', 'main', 0, 'feature/kinou2', 'http://example.com/repo1/pr/2', 2, 'REPO_1');