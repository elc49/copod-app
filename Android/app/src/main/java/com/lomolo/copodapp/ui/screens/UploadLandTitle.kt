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
import com.lomolo.copodapp.state.viewmodels.OnboardingViewModel
import com.lomolo.copodapp.state.viewmodels.UploadingDoc
import com.lomolo.copodapp.ui.common.UploadDocument
import com.lomolo.copodapp.ui.navigation.Navigation
import kotlinx.coroutines.launch

object UploadLandTitleScreenDestination : Navigation {
    override val title = null
    override val route = "register-land-title"
    const val RESUBMIT_ID_ARG = "reSubmit"
    val routeWithArgs = "$route/{$RESUBMIT_ID_ARG}"
}

@Composable
fun UploadLandTitle(
    modifier: Modifier = Modifier,
    onGoBack: () -> Unit,
    onNavigateTo: (String) -> Unit,
    viewModel: OnboardingViewModel,
    isResubmit: Boolean?,
) {
    val image by viewModel.landTitle.collectAsState()
    val landTitle = when (viewModel.uploadingLandDoc) {
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
    val pickLandTitleMedia = rememberLauncherForActivityResult(
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
                viewModel.uploadLandTitle(fileName, stream)
            }
        }
    }
    val savingDoc = viewModel.uploadingLandDoc is UploadingDoc.Loading || viewModel.updatingTitleDeed is UploadingDoc.Loading

    UploadDocument(
        modifier = modifier,
        title = @Composable {
            Column {
                Text(stringResource(R.string.land_title_document))
            }
        },
        image = landTitle,
        newUpload = image.isEmpty(),
        savingDoc = savingDoc,
        onNext = {
            if (image.isNotEmpty()) {
                if (isResubmit == true) {
                    viewModel.updateTitleDeed {
                        onNavigateTo(CreateLandScreenDestination.route)
                    }
                } else onNavigateTo("${UploadGovtIssuedIdScreenDestination.route}/${false}")
            }
        },
        onGoBack = onGoBack,
        onSelectImage = {
            if (!savingDoc) {
                scope.launch {
                    pickLandTitleMedia.launch(
                        PickVisualMediaRequest(
                            ActivityResultContracts.PickVisualMedia.ImageOnly,
                        )
                    )
                }
            }
        },
        buttonText = @Composable {
            Text(
                stringResource(R.string.save),
                style = MaterialTheme.typography.titleMedium,
            )
        })
}