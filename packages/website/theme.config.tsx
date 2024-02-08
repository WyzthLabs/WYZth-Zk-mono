import Footer from "./components/Footer";
import ThemedImage from "./components/ThemedImage";
import { useConfig } from "nextra-theme-docs";
import { useRouter } from "next/router";
import ThemeToggle from "./components/ThemeToggle/ThemeToggle";

export default {
  banner: {
    key: "banner",
    text: (
      <a href="/docs/guides" target="_blank">
        📌 Alpha-3 is here! Get started →
      </a>
    ),
  },
  chat: {
    link: "https://discord.gg/wyzth_zkevmxyz",
  },
  darkMode: false,
  docsRepositoryBase:
    "https://github.com/wyzth_zkevmxyz/wyzth_zkevm-mono/blob/main/packages/website",
  editLink: {
    text: "Edit this page ↗",
  },
  feedback: {
    content: null,
  },
  footer: {
    component: Footer,
  },
  head: () => {
    const { asPath } = useRouter();
    const { frontMatter } = useConfig();
    return (
      <>
        <meta property="og:url" content={`https://wyzth_zkevm.xyz${asPath}`} />
        <meta property="og:title" content={frontMatter.title || "WYzth_zkevm"} />
        <meta
          property="og:description"
          content={
            frontMatter.description ||
            "A decentralized, Ethereum-equivalent ZK-Rollup."
          }
        />
        <meta
          property="og:image"
          content={"/images/WYzth_zkevm_social_media_preview.png"}
        />
        <link rel="icon" href="/images/favicon.svg" />
      </>
    );
  },
  logo: <ThemedImage />,
  navbar: {
    extraContent:
  <>
    <ThemeToggle />
  </>,
  },
  nextThemes: {
    defaultTheme: "light",
  },
  primaryHue: 323,
  project: {
    link: "https://github.com/wyzth_zkevmxyz",
  },
  useNextSeoProps() {
    return {
      titleTemplate: "%s – WYzth_zkevm",
    };
  },
};
