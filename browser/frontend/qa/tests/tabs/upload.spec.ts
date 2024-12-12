import { test, expect } from '@playwright/test';

test('has html lang attribute set', async ({ page }) => {
    await page.goto('http://localhost:8080/');

    // Expect a title "to contain" a substring.
    const htmlElement = page.locator("html");
    await expect(htmlElement).toHaveAttribute("lang", "en");
});
