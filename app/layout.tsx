import 'react-datepicker/dist/react-datepicker.css';
import './globals.css';

export const metadata = {
  title: 'Eclipse Dashboard',
  description:
    'Eclipse Academy admin dashboard for tracking students and teachers, built by ozzyfromspace',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
