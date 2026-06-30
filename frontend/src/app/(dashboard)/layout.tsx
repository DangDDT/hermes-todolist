import { AuthGuard } from '@/components/auth-guard';
import { Header } from '@/components/header';

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <AuthGuard>
      <div className="flex min-h-screen">
        <Header />
        <main className="flex-1 overflow-auto pl-0 lg:pl-0">
          <div className="container mx-auto p-6">{children}</div>
        </main>
      </div>
    </AuthGuard>
  );
}
