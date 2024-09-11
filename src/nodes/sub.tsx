import { createJdmNode } from '@gorules/jdm-editor';
import { MinusOutlined } from '@ant-design/icons';

export const subNode = createJdmNode({
  kind: 'mathSub' as string,
  group: 'math',
  icon: <MinusOutlined />,
  displayName: 'Subtract',
  shortDescription: 'Subtracts 2 numbers',
  inputs: [
    { label: 'a', name: 'a', control: 'text' },
    { label: 'b', name: 'b', control: 'text' },
    { label: 'Key', name: 'key', control: 'text' },
  ]
});