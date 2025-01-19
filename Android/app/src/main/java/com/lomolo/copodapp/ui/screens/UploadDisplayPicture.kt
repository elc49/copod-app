package com.lomolo.copodapp.ui.screens

import android.provider.OpenableColumns
import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.PickVisualMediaRequest
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.layout.Column
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.stringResource
import com.lomolo.copodapp.R
import com.lomolo.copodapp.state.viewmodels.Onboarding
import com.lomolo.copodapp.state.viewmodels.OnboardingViewModel
import com.lomolo.copodapp.state.viewmodels.UploadingDoc
import com.lomolo.copodapp.ui.common.UploadDocument
import com.lomolo.copodapp.ui.navigation.Navigation
import kotlinx.coroutines.launch

object UploadDisplayPictureDestination : Navigation {
    override val title = null
    override val route = "register-display-picture"
}

@Composable
fun UploadDisplayPicture(
    modifier: Modifier = Modifier,
    onGoBack: () -> Unit,
    onNext: (String) -> Unit,
    viewModel: OnboardingViewModel,
) {
    val image by viewModel.displayPicture.collectAsState()
    val displayPicture = when (viewModel.uploadingDp) {
        UploadingDoc.Loading -> {
            R.drawable.loading_img
        }

        UploadingDoc.Success -> {
            if (image.isEmpty()) {
                R.drawable.upload
            } else {
                image
            }
        }

        else -> {
            R.drawable.upload
        }
    }

    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    val pickDisplayMedia = rememberLauncherForActivityResult(
        contract = ActivityResultContracts.PickVisualMedia()
    ) {
        if (it != null) {
            val stream = context.contentResolver.openInputStream(it)
            var fileName = "Unknown"
            context.contentResolver.query(it, null, null, null, null)?.use { cursor ->
                if (cursor.moveToFirst()) {
                    val nameIndex = cursor.getColumnIndex(OpenableColumns.DISPLAY_NAME)
                    if (nameIndex >= 0) {
                        fileName = cursor.getString(nameIndex) ?: "${System.currentTimeMillis()}"
                    }
                }
            }
            if (stream != null) {
                viewModel.uploadDisplayPicture(fileName, stream)
            }
        }
    }

    UploadDocument(modifier = modifier,
        titleSize = "large",
        title = @Composable {
            Column {
                Text(stringResource(R.string.display_picture))
            }
        },
        image = displayPicture,
        newUpload = image.isEmpty(),
        savingDoc = viewModel.uploadingDp is UploadingDoc.Loading || viewModel.onboarding is Onboarding.Loading,
        onNext = {
            if (image.isNotEmpty()) {
                viewModel.createOnboarding {
                    onNext(it)
                }
            }
        },
        onGoBack = onGoBack,
        onSelectImage = {
            if (viewModel.uploadingDp !is UploadingDoc.Loading) {
                scope.launch {
                    pickDisplayMedia.launch(
                        PickVisualMediaRequest(
                            ActivityResultContracts.PickVisualMedia.ImageOnly,
                        )
                    )
                }
            }
        },
        buttonText = @Composable {
            Text(
                stringResource(R.string.continue_text),
                style = MaterialTheme.typography.titleMedium,
            )
        })
}