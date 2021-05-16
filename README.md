# github-action-post-to-hacker-news
GitHub Action for Posting Releases To Hacker News.

## Required Environment Variables
* `HN_USERNAME`: The username to use for Hacker News. 
* `HN_PASSWORD`: The password to use for Hacker News. 
* `HN_TITLE_FORMAT_SPECIFIER`: The [Go format specifier](https://golang.org/pkg/fmt/) to use for formatting the title on the Hacker News post. For example, setting `HN_TITLE_FORMAT_SPECIFIER=Celery v%s Released` would make the title posted `Celery v5.1.0 Released` (presuming the `GITHUB_REF` value was `5.1.0` or (`v5.1.0`, see the *Important Note(s)* section below)).
* `HN_URL_FORMAT_SPECIFIER`: The [Go format specifier](https://golang.org/pkg/fmt/) to use for formatting the url on the Hacker News post. Usage and behavior is identical to that of `HN_TITLE_FORMAT_SPECIFIER` above.
* `HN_TEST_MODE` (Optional): If set to `true` (or `True`), instead of actually posting to Hacker News, simply print out the `title` and `url` that would otherwise be posted. This is useful in the beginning for checking that the `title` and `url` are exactly what you'd expect and that the action is behaving as you'd like. Without setting this, posts to Hacker News can get rate limited pretty easily, so it's nice to get the integration to be exactly what you want before you turn off test mode.

## Important Note(s)
* `GITHUB_REF` will be left trimmed of any `"v"` or `"V"` characters. This makes it easier to support releases of the format `v5.1.0` (for example) and `5.1.0`.
