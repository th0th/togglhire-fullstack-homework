declare module '*.svg' {
  export const ReactComponent: React.FC<React.PropsWithoutRef<JSX.IntrinsicElements['svg']>>;

  const content: any;
  export = content;
}
