'use client';

import { usePathname } from 'next/navigation';
import { ChevronRight } from 'lucide-react';
import { UserMenu } from './user-menu';
import { ThemeToggle } from './theme-toggle';
import { Sidebar } from './sidebar';

function getBreadcrumbs(pathname: string) {
  const segments = pathname.split('/').filter(Boolean);
  const breadcrumbs: { label: string; href: string }[] = [];

  let accumulated = '';
  for (const segment of segments) {
    accumulated += `/${segment}`;
    const label = segment
      .replace(/-/g, ' ')
      .replace(/\b\w/g, (c) => c.toUpperCase());

    // If segment looks like a UUID, show "Task Detail" or "Edit Task" based on path
    if (/^[0-9a-f-]{36}$/i.test(segment)) {
      const parentSegment = segments[segments.indexOf(segment) - 1] || '';
      if (parentSegment === 'tasks') {
        if (segments.includes('edit')) {
          breadcrumbs.push({ label: 'Edit Task', href: accumulated + '/edit' });
        } else {
          breadcrumbs.push({ label: 'Task Detail', href: accumulated });
        }
        continue;
      }
    }

    breadcrumbs.push({ label, href: accumulated });
  }

  return breadcrumbs;
}

export function Header() {
  const pathname = usePathname();
  const breadcrumbs = getBreadcrumbs(pathname);

  return (
    <>
      <Sidebar />
      <header className="sticky top-0 z-30 flex h-14 items-center gap-4 border-b bg-background px-4 lg:px-6">
        <div className="flex items-center gap-2 text-sm text-muted-foreground">
          {breadcrumbs.map((crumb, index) => (
            <span key={crumb.href} className="flex items-center gap-2">
              {index > 0 && <ChevronRight className="h-4 w-4" />}
              {index === breadcrumbs.length - 1 ? (
                <span className="font-medium text-foreground">{crumb.label}</span>
              ) : (
                <span>{crumb.label}</span>
              )}
            </span>
          ))}
        </div>

        <div className="ml-auto flex items-center gap-2">
          <ThemeToggle />
          <UserMenu />
        </div>
      </header>
    </>
  );
}
