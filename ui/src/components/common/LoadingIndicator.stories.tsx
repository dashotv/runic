import React from 'react';

import type { Meta, StoryObj } from '@storybook/react';

import { LoadingIndicator } from './LoadingIndicator';

const meta: Meta<typeof LoadingIndicator> = {
  component: LoadingIndicator,
};

export default meta;
type Story = StoryObj<typeof LoadingIndicator>;

export const Default: Story = {
  render: () => <LoadingIndicator />,
};
