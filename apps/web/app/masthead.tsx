import Link from "next/link";

export default function Masthead({ practice = false }: { practice?: boolean }) {
  return (
    <header className="mast">
      <nav>
        <Link className="brand" href="/">
          <i aria-hidden="true">⌁</i>Incident
        </Link>
        <span className="nav-rule" />
        <Link href="/">Practice library</Link>
        {practice && <span className="crumb">/ Investigation</span>}
      </nav>
    </header>
  );
}
