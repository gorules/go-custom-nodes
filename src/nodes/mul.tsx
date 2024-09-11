import { createJdmNode } from '@gorules/jdm-editor';

export const mulNode = createJdmNode({
  kind: 'mathMul' as string,
  group: 'math',
  icon: '*',
  displayName: 'Multiply',
  shortDescription: 'Multiplies 2 numbers',
  inputs: [
    { label: 'a', name: 'a', control: 'text' },
    { label: 'b', name: 'b', control: 'text' },
    { label: 'Key', name: 'key', control: 'text' },
  ]
});