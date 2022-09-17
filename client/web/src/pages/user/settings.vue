<template>
  <div class="bg-#fff box-border border-rd-4px relative p-20px">
    <NTabs
      default-value="profile"
      size="large"
      animated
    >
      <NTabPane name="profile" tab="个人资料">
        <NForm>
          <NFormItemRow label="头像" path="avatar">
            <NUpload
              action="https://www.mocky.io/v2/5e4bafc63100007100d8b70f"
              :file-list="[]"
              list-type="image-card"
            />
          </NFormItemRow>
          <NFormItemRow label="昵称" path="nickname" required>
            <NInput />
          </NFormItemRow>
          <NFormItemRow label="个人简介" path="headline">
            <NInput />
          </NFormItemRow>
        </NForm>
        <div>
          <NButton type="primary" secondary strong>
            保存修改
          </NButton>
        </div>
      </NTabPane>
    </NTabs>
  </div>
</template>

