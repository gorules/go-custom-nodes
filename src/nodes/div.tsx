import { createJdmNode } from '@gorules/jdm-editor';

export const divNode = createJdmNode({
  kind: 'mathDiv' as string,
  group: 'math',
  icon: '/',
  displayName: 'Divide',
  shortDescription: 'Divides 2 numbers',
  inputs: [
    { label: 'a', name: 'a', control: 'text' },
    { label: 'b', name: 'b', control: 'text' },
    { label: 'Key', name: 'key', control: 'text' },
  ]
});